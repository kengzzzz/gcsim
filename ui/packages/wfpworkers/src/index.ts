import { Router } from "itty-router";
import { handleAssets } from "./assets";
import { handleShare, handleView } from "./share";
import { handleWasm } from "./wasm";

const router = Router();

router.post("/api/share", handleShare);
router.get("/api/share/:key", handleView);
router.get("/api/assets/*", handleAssets);
router.get("/api/wasm/*", handleWasm);

addEventListener("fetch", (event) => {
  event.respondWith(router.handle(event.request, event));
});
