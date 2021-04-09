# ipify-api

*A Simple Public IP Address API*

This is a fork of [the original repository](https://github.com/rdegges/ipify-api) with some modifications. A few notible ones are:

- Obtaining the IP address of the connecting client directly, if not running behind a proxy.
- XML format available on output, if using the correct URL.
- Precompiled binaries available for download.
- Example systemd service provided.
- Documentation for deploying on a non-Heroku host.
- Documentation of URL endpoints within the repository itself.

This repository contains the source code for [ipify](https://www.ipify.org), one
of the largest and most popular IP address API services on the internet. ipify serves over 30 billion requests per month!


## What does ipify do?

Have you ever needed to pragmatically get your public IP address? This is quite common for developers provisioning cloud servers, for instance, where you might be creating servers and running bootstrapping software on them without access to server metadata.

Being able to quickly and reliably get access to your public IP address is essential for configuring DNS, managing external services, and a number of other operationally related tasks.

In general, there are a number of uses for public IP address information.


## What is ipify?

ipify is a free API service anyone can use to get their public IP address. It is highly reliable (built on top of [Heroku](https://www.heroku.com/)) and fast. Typical response times (server side) are between 1ms and 10ms.

ipify is also fully funded -- it's been running for years and isn't going anywhere. The people behind ipify cover all expenses and maintenance, so you can feel safe integrating with it knowing it won't be disappearing.

If you'd like to use ipify in your application, no permission is needed. You can immediately start using the service without any restrictions. Simply visit our [public website](https://www.ipify.org) for more information.


## What is this project?

This project is the source code that powers the ipify service. ipify is written in the Go programming language for speed and efficiency purposes. You can read an [article](https://www.rdegges.com/2018/to-30-billion-and-beyond/) written by ipify's creator, [Randall Degges](https://twitter.com/rdegges), if you'd like more information.

If you'd like to contribute to ipify's development, you can do so here. Pull requests are encouraged.

Finally, if you'd like to deploy your own instance of ipify, you can easily do so. Compiling this project will produce a binary that is designed to be run on Heroku. With minor modification, ipify can be ran on any web hosting platform.


## Building ipify

To develop and build ipify, you'll need to have the Go programming language setup on your computer. If you don't, you can read more about it [here.](https://golang.org/)

Once you have Go installed, you'll need to clone this project.

```bash
$ git clone https://github.com/tech10/ipify-api.git
```

To build the project, change to the project directory and run:

```bash
$ go build
```

This will create the `ipify-api` binary in the current directory that you can use for testing or deployment, if desired. This will not be built as a static binary by default at the moment.

To build a static binary, execute the following script, which has been tested on Linux and is known to build a static binary correctly.

```bash
$ ./build.sh
```


## Deploying ipify

If you'd like to deploy your own version of ipify to Heroku, you can do so easily by clicking the button below. This will take you to Heroku and let you instantly provision your own copy of the ipify service.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)


### Deploying on alternate platforms

If you are not using Heroku or you'd simply like more customization for your build, here are some basic instructions on how to get started.

To deploy this application, you can either build the server yourself, or download one of the pre-compiled binaries available on the releases page. You can then use your operating systems daemon manager to start the server, or you can manually start it yourself. Whether or not you fork it into the background will then be up to you.

By default, the program will listen for incoming requests on port 3000. To change this, run the program with the PORT variable modified. For instance, to listen on port 8080, run the program in the following way.

```bash
$ PORT=8080 ./ipify-api
```

Multiple methods are supported in order to obtain your IP address. You can proxy to the server and use the X-Forwarded-For header through the use of another web server, or you can obtain a direct connection. When you proxy your requests, you have the option of using https support. Otherwise, the program will only listen using the http protocol. Precise instructions on how to set up your web server to proxy requests to the ipify-api server are beyond the scope of this document, but you can easily obtain this information with a simple search in your favorite search engine.

## Valid endpoints for IP address format

Most of these are documented on the ipify website itself. However, this particular server has four output options available. They are json, jsonp, xml, and text. If no format parameter is provided, the default will be text.

The domain provided in example URL's is not a production domain, it's only there for purposes of documenting available endpoints. The default port is used to demonstrate obtaining your IP address if you aren't using a proxy, and wish to use a direct URL to the server itself.

- Retrieve json output: `http://example.org:3000/?format=json`
- Retrieve jsonp output: `http://example.org:3000/?format=jsonp`
- Retrieve jsonp output with callback: `http://example.org:3000/?format=jsonp&callback=getip
- Retrieve xml output: `http://example.org:3000/?format=xml`
- Retrieve text output:
    - `http://example.org:3000`
    - `http://example.org:3000/?format=text`
    - `http://example.org:3000/?format=anythingAtAll`

Anything in the format parameter that isn't recognized will automatically default to text output of the IP address.


## Obtaining IPV4 and IPV6 addresses

Since the server will automatically listen on all IPV4 and IPV6 addresses, you can obtain your IPV4 address, IPV6 address, or which ever resolves first if you are using a domain. You can use a URL in the following ways:

- A direct IPV4 address.
- A direct IPV6 address.
- A domain with A and AAA records for resolving to IPV4 and IPV6 addresses.
- A domain with different A and AAA records for resolving to only an IPV4 or IPV6 address.

Precise options for setting this up are beyond the scope of this document. Please refer to your server host or domain registrar for information on how this can be achieved.


## Questions

If you have any questions on how to use this server, any issues are welcome. Please try and limit your questions to downloading and running the server, rather than configuring it to work with your server host directly. I will try and answer any questions that I can, however.

Good luck on providing IP addresses to users, or obtaining your own!
