export {};

declare global {
  const ASSETS_R2: R2Bucket; //bucket
  const WASM_R2: R2Bucket; //bucket
  const WFPSIM_KV: KVNamespace; //kv
  const PREVIEW_ENDPOINT: string;
  const AUTH_KEY: string;
}
