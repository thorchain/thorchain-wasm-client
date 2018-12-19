import { broadcastTxCommit } from '../request'
import { Runner } from '../runner/Runner'
import { IKey } from './IKey'
import { ITxContext } from './ITxContext'

export class Client {
  constructor(private runner: Runner, private uri: string) {}

  public broadcastTx(signedTx: string) {
    return broadcastTxCommit(this.uri, signedTx)
  }

  public helloWorld(encoded: string) {
    return this.runner.invoke<string>('helloWorld', encoded)
  }

  public async createKey() {
    const response = await this.runner.invoke<string>('createKey')
    return JSON.parse(response) as IKey
  }

  public async getPubAndAddrFromPrivKey(privStr: string) {
    const response = await this.runner.invoke<string>('getPubAndAddrFromPrivKey', privStr)
    return JSON.parse(response) as IKey
  }

  public signSendTx(txContext: ITxContext, from: string, to: string, amount: string) {
    return this.runner.invoke<string>('signSendTx', txContext, from, to, amount)
  }

  public async signAndBroadcastSendTx(txContext: ITxContext, from: string, to: string, amount: string) {
    const signedTx = await this.signSendTx(txContext, from, to, amount)
    return this.broadcastTx(signedTx)
  }

  public signClpTradeTx(txContext: ITxContext, from: string, fromTicker: string, toTicker: string, fromAmount: number) {
    return this.runner.invoke<string>('signClpTradeTx', txContext, from, fromTicker, toTicker, fromAmount)
  }

  public async signAndBroadcastClpTradeTx(txContext: ITxContext, from: string, fromTicker: string, toTicker: string,
    fromAmount: number) {
    const signedTx = await this.signClpTradeTx(txContext, from, fromTicker, toTicker, fromAmount)
    return this.broadcastTx(signedTx)
  }

  public signExchangeCreateLimitOrderTx(txContext: ITxContext, sender: string, kind: 'buy'|'sell', amount: string,
    price: string, expiresAt: string) {
    return this.runner.invoke<string>('signExchangeCreateLimitOrderTx', txContext, sender, kind, amount, price,
      expiresAt)
  }

  public async signAndBroadcastExchangeCreateLimitOrderTx(txContext: ITxContext, sender: string, kind: 'buy'|'sell',
    amount: string, price: string, expiresAt: string) {
    const signedTx = await this.signExchangeCreateLimitOrderTx(txContext, sender, kind, amount, price, expiresAt)
    return this.broadcastTx(signedTx)
  }
}
