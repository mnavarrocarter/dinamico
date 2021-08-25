<h1 align="center">
Dinamico
</h1>
<p align="center">
The hero your home lab needs
</p>
<hr/>

Dinamico is a utility written in Go to update your dynamic IP records across
a variety of providers. It is meant to be a replacement for tools like `ddclient`
and others. 

## Features:

- [x] Dead simple configuration format
- [ ] Support for the most popular dynamic DNS providers
- [ ] REST api to add records programmatically to your watch list
- [ ] HTML UI to add records from the comfort of your web browser
- [ ] Graceful shutdown writing state to disk

## Usage

```
dinamico -log -state=domains.txt
```

Where `domains.txt` is a file that contains the records to update.

## Records

Every record is represented by a URI. The provider is specified in the `Scheme`. 
Authentication credentials are specified in the `UserInfo` part of the uri, and 
the hostname to update is specified in the `Host` part.

For instance, to add a watcher for a Google Domains Dynamic Record:

```
google-domains://username:password@my.hostname.com
```

## Supported Providers

At the moment, support for providers in very limited. I am working into reaching v1.0
that will include support for all the providers currently supported by `ddclient`.

### Google Domains

