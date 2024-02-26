# Play around with the CRUDE BlockChain
## Create a blog post by Alice
```blogd tx blog create-post hello world --from alice --chain-id blog```
## Show a blog post
```blogd q blog show-post 0```
## Create a blog post by Bob
```blogd tx blog create-post foo bar --from bob --chain-id blog```
## List all blog posts
```blogd q blog list-post```
## Update a blog post
```blogd tx blog update-post hello cosmos 0 --from alice --chain-id blog```

```blogd q blog show-post 0```
## Filter blog posts
Display the posts whose title contains `keyword` 
```blogd q blog filter-post keyword```
## Delete a blog post
```blogd tx blog delete-post 0 --from alice --chain-id blog```
## Delete a blog post unsuccessfully
```blogd tx blog delete-post 1 --from alice --chain-id blog```

# Breaking the consensus
Please refer to the `consensus-breaking` branch.

A consensus-breaking change is a change to the blockchain network that invalidates some prior agreements or "game rules". For example, a change in the formula for reward calculation can be seen as a consensus-breaking change. Other examples include a change in the format of APIs or messages, a change in the underlying communication protocol (TCP, UDP, etc), and a change in the signature scheme. Other people who want to remain in the network must update their end, or otherwise, a fork may happen.

In this branch, I cancelled the identity check for `UpdatePost`. This means all posts become wiki-style posts, and for example, Bob can update Alice's posts. The prior agreement that everyone exclusively owns their posts is broken.
