import { Client } from './client/Client'
import { IRuntime } from './interfaces/IRuntime'
import { Runner } from './runner/Runner'

export default async function (
  runtime: IRuntime,
  nodeUri: string = 'http://localhost:26657',
  wasmUri: string = '/thorchain_wasm_client.wasm',
) {
  const runner = new Runner(runtime, wasmUri)
  await runner.init()

  const client = new Client(runner, nodeUri)

  return {
    client,
    runner,
  }
}
