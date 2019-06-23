# Quick FileWallet reference

## Local structure
* One file can hold multiple wallets, each wallet is identified by a name and encrypted with a masterpass
* Wallets are kept logically separated within the file
* Every wallet name, key and value is encoded with the masterpass and nonces, to ensure data privacy

## Wallet structure
* Relevant wallet keys are
    1. WalletInfoKey - Contains wallet name
    2. BlockReferenceKey - Contains a list of known block IDs 
    3. LastBlockReferenceKey - Contains the ID of the last know block
    4. OutputReferenceKey - Contains a list known output IDs
    5. OutputTypeReferenceKey - Contains a list of known output types
    6. UnspentOutputReferenceKey - Contains a list of known unspent output IDs



    * Single object keys:
        1. "**Out-**" + **outputID**    - Single output saved as 
        2. "**Blk-**" + **blockID**     - Single block header
        3. "**Typ-**" + **outputType**  - List of output IDs referring to outputs of the given type

* How are **_IDs_** calculated
    * **outputID**      - **byte(** blockHash **)** + **byte(** output.LocalIndex **)** ; 
    * **blockID**       - **byte(** blockHash **)**
    * **outputType**    - **string**