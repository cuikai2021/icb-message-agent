# icb-message-agent

This agent is used for sending a stream of messages to customers when the tools are running.

**Principle**: All messages visible to the customers need to be **reviewed** .

The messages sent by the tools should  be in the style of format string(message template) and parameter list.

Before using a message template, it should be added to the `templates` .



## message templates

All legal message templates are stored in the file under the folder `templates` .

Each message template occupies a separate line，

The agent Judging whether the message to be sent is legal based on the policy of exact string matching.




## gRPC protocol

The messages are finally sent out by the Ginkgo server,
so the agent provides the definition of the gRPC interface in `proto`.

If a new programming language needs to be introduced,
the developer needs to generate the corresponding client code and complete the encapsulation of the agent logic.

The only Interface in the protocol is ``SendMessage``.


```protobuf
rpc SendMessage (SendMessageRequest) returns (google.protobuf.Empty) {
}
```



## agent

The main logic of the agent

- Check the legitimacy of the message by message templates

- Send the message to customer by Ginkgo gRPC



### message legal checker

The policy of the checker is exact string matching, messages which that fail the legality check are discarded.



### gRPC sender

What we need to be done in gRPC sender are wrap the Request body and send gRPC request to Ginkgo Server.

Things that need special clarification

- The Ginkgo Server Address
  The environment variable `DEPLOY_MODEL` is used to identify the current deployment environment.
  If the value is `staging`, use `ginkgo.internal.icbench.com:1443` as server address; otherwise use `ginkgo.grpc.icbench.com:1443`.

- The Running taskId
  In the request of `SendMessage`, the value for `taskId` can get from environment variable `ICB_RUN_ID`.



### batch sender(Optional)

To reduce the number of server side requests for sending messages, agent could use client side cache for sending message in batches.


