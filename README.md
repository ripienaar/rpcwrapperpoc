# What?

A tool that generates convenient Go APIs from Choria Agent DDLs.

The APIs are inspired by the Google Cloud APIs and are composed using chained methods, this allow us to expand the options in future while offering comfortable IDE integration.

# Status?

I am happy with the generated API and the generator CLI is functional, this is ready for wider testing with an eye on incorporation into the next Choria release.

# API?

The API is around a chaining mode, if you wanted to do the following CLI command:

```
mco rpc puppet disable message="testing, 1, 2, 3" -F customer=acme
```

In go this should be the minimal needed code:

```golang
p, err := puppet.New()
panicIfErr(err)

res, err = p.InBatches(10, 30).FactFilter("customer=acme"). \\ these are client level settings that persist between
                                                            \\ invocations of any agent
  Disable().                                                \\ if there are mandatory inputs to the action `disable`
                                                            \\ you would pass them as arguments here, optional ones \\ are chained
  Message("testing, 1, 2, 3").Do(ctx)

panicIfErr(err)
```

This is all that's needed, it will use the users default config, certs etc and disable puppet on `customer=acme` with a message, this is basically:

Processing the results is like this:

```golang
res.EachOutput(func(r *puppet.DisableResult) {
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

```golang
// channel to receive outputs
c := make(chan *puppet.DisableOutput, 1000)

// res is still a result but its output are empty
res, err = p.InBatches(10, 30).FactFilter("customer=acme").Disable().Results(c).Do(ctx)
```

This will pass the replies into the channel as they are received and you can process them in another go routine.

