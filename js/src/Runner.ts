import { config } from './config'
import { IRuntime } from './interfaces/IRuntime'

export class Runner {
  private instance: WebAssembly.Instance

  constructor(private exec: IRuntime, private binaryUri: string) {}

  public async init () {
    const result = await this.instantiateStreaming(fetch(this.binaryUri), this.exec.importObject)
    this.instance = result.instance
    this.exec.run(this.instance)
  }

  public invoke<T>(funcName: string, ...args: any[]): Promise<T> {
    return new Promise((resolve, reject) => {
      global[config.namespace][funcName](...args, (err: string, data: T) => {
        if (err) {
          console.log('got err', err)
          reject(err)
        } else {
          resolve(data)
        }
      })
    })
  }

  // Loading in this way doesn't require .wasm files to be served with
  // a special content type.
  private async instantiateStreaming(resp: Promise<Response>, importObject: WebAssembly.Imports) {
    const source = await (await resp).arrayBuffer()
    return await WebAssembly.instantiate(source, importObject)
  }
}