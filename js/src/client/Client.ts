import { broadcast } from '../request'
import { Runner } from '../runner/Runner'
import { IKey } from './IKey'
import { ITxContext } from './ITxContext'

export class Client {
  constructor(private runner: Runner, private uri: string) {}

  public helloWorld(encoded: string) {
    return this.runner.invoke<string>('helloWorld', encoded)
  }

  public createKey() {
    return this.runner.invoke<IKey>('createKey')
  }

  public signTx(txContext: ITxContext, from: string, to: string, amount: string) {
    return this.runner.invoke<string>('signTx', txContext, from, to, amount)
  }

  public broadcastTx(signedTx: string) {
    return broadcast(this.uri, signedTx)
  }

  public async signAndBroadcastTx(txContext: ITxContext, from: string, to: string, amount: string) {
    const signedTx = await this.signTx(txContext, from, to, amount)
    return this.broadcastTx(signedTx)
  }
}
