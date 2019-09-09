# What?

A POC to see what a wrapper API might look like around a choria rpc agent.

The main idea is that one might run `choria tool generate client puppet` and this will, based on the puppet DDL` produce a API client that looks a bit like this.

The code will be generated and based on the DDL.

# Status?

It's kind of working and I like the design, no generators or anything, this is about playing with the eventual outcome API.

# API?

The API is around a chaining mode, this should be the minimal:

```golang
p, err := puppet.New()
panicIfErr(err)

res, err = p.InBatches(10, 30).FactFilter("customer=acme"). \\ these are client level settings that persist between
                                                            \\ invocations of any agent
  Disable().                                                \\ if there are mandatory inputs to the action `disable`
                                                            \\ you'd pass them as arguments here, optional ones are chained
  Message("testing, 1, 2, 3).
  Do(ctx)
  
panicIfErr(err)
```

This is all that's needed, it will use the users default config, certs etc and disable puppet on `customer=acme` with a message, this is basically:

```
mco rpc puppet disable message="testing, 1, 2, 3"
```

Processing the results is like this:

```golang
res.EachResult(func(r *puppet.DisableResult) {
  // the disable action has 2 outputs "status" and "enabled"
  // data types are either correct if the DDL defined it else interface{},
  // here its string and boolean
  fmt.Printf("%s: status: %s enabled: %v\n", r.ResultDetails().Sender(), r.Status(), r.Enabled())
  
  // status code and message is accessible
  r.ResultDetails().StatusMessage()
  r.ResultDetails().StatusCode()
  
  // you can also get a map[string]interface{}
  r.HashMap()
  
  // or just all in json which would be more detailed including sender etc
  r.JSON()
})
```

or for faster processing this:

```
c := make(chan *puppet.DisableResult, 1000)

res, err = p.InBatches(10, 30).FactFilter("customer=acme").Disable().Results(c).Do(ctx)
```

This will pass the replies into the channel as they are received and you can process them in another go routine.

