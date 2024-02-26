# System Setup
Suppose in the network there is a broadcaster `B` and some listners `L`. `B` maintains the following:
1. A self-incrementing ID for transactions accepted and signed (for which a 200 code is returned). This ID can globally identify a transaction.
1. A logical bi-directional channel, `C_L`, for each `L`. 
1. The highest consecutive ID from the initial value, `ID_L`, responded by each `L`. This mechanism is similar to the ACK number applied in TCP.
1. A clock that keeps track of the time `T`.
1. The wall-clock timestamp, `TS_L`, of the most recent response from `L`.
1. A list of transactions `TX` whose `ID_TX` is greater than `ID_L` for some `L`. These transactions are promised but undelivered. The master copy of this list should be in the disk or some permanent storage. A new transaction should be fully added to the master copy before the broadcaster returns 200. Another copy can be stored in the memory and regularly pulled from the master copy. This mechanism prevents lossing undelivered transactions if the broadcaster fails.

# Listener
1. Maintains a variable `ID_L`.
1. On start-up or recovery from crash, sends `ID_L` to the broadcaster `B`.
    1. If a transaction of `ID_TX` higher than or equal to `ID_L` is received, updates `ID_L` according to its definition. Sends `ID_L` to `B`.
    1. If no transactions from `B` are received within some time, say `a` seconds, re-sends `ID_L`. Goes to the previous step.
1. Upon receipt of a transaction, updates `ID_L` accordingly. Send `ID_L`.

# Broadcaster
Normal State:

    1. Accepts new transaction requests (returns 200 OK) if the size of the list of undelivered transactions is below some threshold, say `b`. Send new transactions to every `L`.
    2. On receipt of `ID_L` for each `L`, if `ID_L` is less than some `ID_TX` in the undelivered list, then send such transactions to `L` immediately.
    3. If the same `ID_L` is received for `c` times or is received `d` seconds ago, transits to the abnormal state. Marks this `L` as faulty.
    4. If no messages have been received from some `L` for `e` seconds and `ID_L` is less than some `ID_TX` in the undelivered list, transits to the abnormal state. Marks this `L` as faulty.

Abnormal State:

    1. Rejects every new transaction request (returns 503 Service Unavailable).
    2. Sends existing undelivered transactions to each non-faulty `L` as in the normal state.
    3. Handles `ID_L` and new faulty `L` as in the normal state.
    4. A faulty `L` becomes non-faulty if it has been active sending responses for `f` seconds or `g` times.
    5. Transits to the normal state if there are no faulty `L`.

The parameters, `a, b, c, d, e, f, g` can be fine-tuned to meet the time and probability requirements.