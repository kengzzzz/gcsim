# 1. build the html assets and publish to pages; note that we need to remove the wasm
# because that's too big
yarn workspace @gcsim/wfpsim build
rm ./packages/wfp/dist/main.wasm
wrangler pages deploy ./packages/wfp/dist --project-name="wfpsim" --branch="main"

# 2. upload the wasm to r2
# clean out any exists
rm ./packages/wfp/wasm/*.wasm
# build new
GCSIM_SHARE_KEY="" && yarn workspace @gcsim/wfpsim build:wasm:web
# rename the wasm based on current git hash
hash=$(git rev-parse HEAD)
echo "hash is: $hash"
cp ./packages/wfp/public/main.wasm "./packages/wfp/wasm/${hash}.wasm"
# use rclone to sync so we get rid of any old ones
rclone sync ./packages/wfp/wasm/ wfpr2:wfpsim/ --progress --config ~/.dotfiles/rclone/rclone.cfg

# 3. publish worker scripts
cd ./packages/wfpworkers
wrangler deploy