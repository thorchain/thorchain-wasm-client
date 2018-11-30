export interface ITxContext {
    priv_key: string
    account_number: number
    sequence: number
    gas: number
    chain_id: string
    memo: string
    fee: string
}