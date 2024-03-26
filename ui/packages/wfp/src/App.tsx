import { Button, Tab, Tabs } from "@blueprintjs/core";
import { Executor, ExecutorSupplier, WasmExecutor } from "@gcsim/executors";
import { SimResults } from "@gcsim/types";
import { useLocalStorage } from "@gcsim/utils";
import axios from "axios";
import { throttle } from "lodash-es";
import { useEffect, useRef, useState } from "react";
import { Sample } from "./Components/Sample/Sample";
import { Simulator } from "./Components/Simulator/Simulator";
import { Viewer, WebViewer } from "./Components/Viewer/Viewer";

const minWorkers = 1;
const maxWorkers = 30;

let exec: WasmExecutor | undefined;

function wasmLocation() {
  if (import.meta.env.PROD) {
    return "/api/wasm/" + import.meta.env.VITE_GIT_COMMIT_HASH + ".wasm";
  }
  return "/main.wasm";
}

type simRunResult = {
  data: SimResults;
  hash: string;
} | null;

const App = ({}) => {
  const [workers, setWorkers] = useLocalStorage<number>("wasm-num-workers", 5);
  const [tabId, setTabId] = useState("simulator");
  const [results, setResults] = useState<simRunResult>(null);
  const [webResult, setWebResult] = useState<SimResults | null>(null);
  const [err, setError] = useState<string>("");

  let key = "";
  const res =
    /sh\/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$/.exec(
      window.location.pathname
    );
  if (res && res.length >= 1) {
    key = res[1];
  }

  useEffect(() => {
    if (key !== "") {
      axios
        .get("/api/share/" + key, { timeout: 30000 })
        .then((resp) => {
          setWebResult(resp.data);
          console.log(resp.data);
        })
        .catch((e) => {
          setError(e.message);
        });
    } else {
      setWebResult(null);
    }
  }, [key, setResults]);

  const supplier = useRef<ExecutorSupplier<WasmExecutor>>(() => {
    if (exec == null) {
      exec = new WasmExecutor(wasmLocation());
      exec.setWorkerCount(workers);
    }
    return exec;
  });

  if (key != "") {
    if (webResult === null) {
      return <div>Loading share please wait</div>;
    }

    return (
      <div className="flex flex-col flex-grow w-full pb-6">
        <div className="px-2 py-4 w-full 2xl:mx-auto 2xl:container">
          <WebViewer data={webResult} />
        </div>
      </div>
    );
  }

  const runSim = (cfg: string) => {
    console.log("starting run");
    setResults(null);
    setError("");

    const updateResult = throttle(
      (res: SimResults, hash: string) => {
        if (tabId !== "results") {
          setTabId("results");
        }
        setResults({ data: res, hash: hash });
      },
      100,
      { leading: true, trailing: true }
    );

    supplier
      .current()
      .run(cfg, updateResult)
      .catch((err) => {
        console.log("problems :(", err);
        setError(err);
      });
  };

  const updateWorkers = (num: number) => {
    num = Math.min(Math.max(num, minWorkers), maxWorkers);
    setWorkers(num);
    supplier.current().setWorkerCount(num);
  };

  const tabs: { [k: string]: React.ReactNode } = {
    simulator: <Simulator exec={supplier.current} run={runSim} />,
    results:
      results === null ? (
        <></>
      ) : (
        <Viewer
          data={results.data}
          hash={results.hash}
          exec={supplier.current}
        />
      ),
    sample:
      results === null ? (
        <></>
      ) : (
        <Sample data={results.data} exec={supplier.current} />
      ),
  };

  if (err !== "") {
    return (
      <div>
        oops something went wrong:
        <br />
        {JSON.stringify(err)}
        <br />
        <Button
          icon="refresh"
          onClick={() => {
            supplier.current().cancel();
            setError("");
            setResults(null);
          }}
          intent="primary"
        >
          Reload
        </Button>
      </div>
    );
  }

  return (
    <div className="flex flex-col flex-grow w-full pb-6">
      <div className="px-2 py-4 w-full 2xl:mx-auto 2xl:container">
        <Tabs selectedTabId={tabId} onChange={(s) => setTabId(s as string)}>
          <Tab
            id="simulator"
            className="focus:outline-none"
            title="Simulator"
          ></Tab>
          <Tab
            id="results"
            className="focus:outline-none"
            title="Results"
            disabled={results === null}
          ></Tab>
          <Tab
            id="sample"
            className="focus:outline-none"
            title="Sample"
            disabled={results === null}
          ></Tab>
          <Tabs.Expander />
        </Tabs>
      </div>
      {tabs[tabId]}
    </div>
  );
};

export function useRunningState(exec: ExecutorSupplier<Executor>): boolean {
  const [isRunning, setRunning] = useState(true);

  useEffect(() => {
    const check = setInterval(() => {
      setRunning(exec().running());
    }, 100 - 50);
    return () => clearInterval(check);
  }, [exec]);

  return isRunning;
}

export default App;
