[![Hex.pm](https://img.shields.io/hexpm/v/restc.svg?maxAge=2592000)](https://hex.pm/packages/restc) restclient -- An erlang REST Client library
====================================

## DESCRIPTION

restclient is a library to help with consuming RESTful web services. It supports
encoding and decoding JSON, Percent and XML and comes with a convenience
function for working with urls and query parameters.

## USAGE

Include restclient as a rebar dependency with:

	{deps, [{restc, ".*", {git, "git://github.com/kivra/restclient.git", {tag, "0.5.0"}}}]}.

You have to start inets before using the client and if you want to use https make sure to start ssl before.
Then you can use the client as:

``` erlang
Erlang/OTP 19 [erts-8.2] [source] [64-bit] [smp:8:8] [async-threads:0] [kernel-poll:false]

Eshell V8.2  (abort with ^G)
1> application:ensure_all_started(restc).
{ok,[idna,mimerl,certifi,ssl_verify_fun,metrics,hackney,
     mochiweb_util,restc]}

2> restc:request(get, "http://148.100.98.30:3000/api/org.acme.sample.Sensor/1").
{ok,200,
    [{<<"Vary">>,<<"Origin, Accept-Encoding">>},
     {<<"Access-Control-Allow-Credentials">>,<<"true">>},
     {<<"X-XSS-Protection">>,<<"1; mode=block">>},
     {<<"X-Frame-Options">>,<<"DENY">>},
     {<<"X-Download-Options">>,<<"noopen">>},
     {<<"X-Content-Type-Options">>,<<"nosniff">>},
     {<<"Content-Type">>,<<"application/json; charset=utf-8">>},
     {<<"Content-Length">>,<<"129">>},
     {<<"ETag">>,<<"W/\"81-qsIqKh9IFJOZdCGV1h71JAfJyOs\"">>},
     {<<"Date">>,<<"Wed, 14 Nov 2018 14:28:03 GMT">>},
     {<<"Connection">>,<<"keep-alive">>}],
    [{<<"$class">>,<<"org.acme.sample.Sensor">>},
     {<<"teamID">>,<<"1">>},
     {<<"teamName">>,<<"Bluebarricade">>},
     {<<"sensorTemp">>,0},
     {<<"thermostatTemp">>,0},
     {<<"recommendation">>,<<>>}]}
3> restc:request(get, "https://api.github.com/herp-derp-404", [200]).
{error,404,
       [{<<"Server">>,<<"GitHub.com">>},
        {<<"Date">>,<<"Thu, 11 May 2017 07:37:27 GMT">>},
        {<<"Content-Type">>,<<"application/json; charset=utf-8">>},
        {<<"Content-Length">>,<<"77">>},
        {<<"Status">>,<<"404 Not Found">>},
        {<<"X-RateLimit-Limit">>,<<"60">>},
        {<<"X-RateLimit-Remaining">>,<<"56">>},
        {<<"X-RateLimit-Reset">>,<<"1494491776">>},
        {<<"X-GitHub-Media-Type">>,<<"github.v3">>},
        {<<"Access-Control-Expose-Headers">>,
         <<"ETag, Link, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit"...>>},
        {<<"Access-Control-Allow-Origin">>,<<"*">>},
        {<<"Content-Security-Policy">>,<<"default-src 'none'">>},
        {<<"Strict-Transport-Security">>,
         <<"max-age=31536000; includeSubdomains; preload">>},
        {<<"X-Content-Type-Options">>,<<"nosniff">>},
        {<<"X-Frame-Options">>,<<"deny">>},
        {<<"X-XSS-Protection">>,<<"1; mode=block">>},
        {<<"X-GitHub-Request-Id">>,
         <<"8C1D:5C90:54F34B8:6C6FF4D:59"...>>}],
       [{<<"message">>,<<"Not Found">>},
        {<<"documentation_url">>,
         <<"https://developer.github.com/v3">>}]}

```

There's also convenience functions for working with urls and query string:

``` erlang
	7> restc:construct_url("http://www.example.com/te", "res/res1/res2", [{"q1", "qval1"}, {"q2", "qval2"}]).
	"http://www.example.com/te/res/res1/res2?q1=qval1&q2=qval2"
```

## License
The KIVRA restclient library uses an [MIT license](http://en.wikipedia.org/wiki/MIT_License). So go ahead and do what
you want!

Lots of fun!
