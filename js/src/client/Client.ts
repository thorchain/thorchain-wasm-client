import { broadcastTxCommit } from '../request'
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
    return broadcastTxCommit(this.uri, signedTx)
  }

// TODO may implement decoding in this module to no longer require LCD to run
  // public async getAccount (address: string) {
  //   const res = await abciQuery(this.uri, {
  //     data: Buffer.from(address, 'utf8').toString('hex'),
  //     height: '0',
  //     path: '/store/acc/key',
  //     trusted: true,
  //   })
  //   // tslint:disable-next-line:no-console
  //   console.log('get account res', res)
  //   // const accountResp = await this.bridge.decodeAccount(JSON.stringify(resp))
  //   // const account = JSON.parse(accountResp)
  //   // return account
  // }

  public async signAndBroadcastTx(txContext: ITxContext, from: string, to: string, amount: string) {
    const signedTx = await this.signTx(txContext, from, to, amount)
    return this.broadcastTx(signedTx)
  }
}
