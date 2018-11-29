export interface IRuntime {
  importObject: WebAssembly.Imports
  run(instance: WebAssembly.Instance): Promise<any>
}