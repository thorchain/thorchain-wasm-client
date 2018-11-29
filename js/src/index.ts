import { Client } from './Client'
import { IRuntime } from './interfaces/IRuntime'
import { Runner } from './Runner'

export default async function (
  runtime: IRuntime,
  wasmUri: string = '/thorchain_wasm_client.wasm',
) {
  const runner = new Runner(runtime, wasmUri)
  await runner.init()

  const client = new Client(runner)

  return {
    client,
    runner,
  }
}
