// Package bridge provides a TCP tunnel to expose services between local networks which
// otherwise would require a VPN to be configured.
//
// As this is a pure userland bridge, no need for admin rights is required.
//
// As a drawback it only works by having a internet facing host which will act as the bridge.
package bridge
