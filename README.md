# Sticker


Sticker is an [Armor](https://github.com/jondot/armor) based reverse-proxy, that
will use a `backend` and an `add_headers` and `remove_headers` configuration,
in order to stick or remove headers onto a web host.

This is needed when you don't have a control over a service, but need to
stick a few headers on to it.


### Examples

You can put Sticker on a free Heroku instance, or an AWS micro instance, and
then handle some of these use cases:

* Stick arbitrary, custom headers on S3 resources
* Fix bad RSS feeds (bad xml/json/atom headers)
* Fix/force control of bad caching with resources (proper datetime)
* Set caching policies on external resources you don't control, infront of CDN
you do control (forcefully CDN another party's resources)

And so on. Mostly, fixing other people's mistakes, and lack of features :)



# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2015 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



