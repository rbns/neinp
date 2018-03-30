/*Package neinp is a toolkit to implement 9p servers.

A 9p filesystem is created by implementing the P2000 interface, which is then used as argument
for NewServer. Server can then use an io.ReadWriter supplied to the Serve method
handle requests using the aforementioned P2000 implementer.

See https://github.com/rbns/rssfs for an example.

NB: This isn't really polished yet, things may break.*/
package neinp
