let _renderDebug = false
export const toggleRenderDebug = () => {
  _renderDebug = !_renderDebug
}
export const getRenderDebug = () => _renderDebug

export const shim = (routes: any, platformWrapper: any, isModal: boolean, isLoggedOut: boolean) => {
  return Object.keys(routes).reduce((map, route) => {
    let _cached = null

    map[route] = {
      ...routes[route],
      // only wrap if it uses getScreen originally, else let screen be special (sub navs in desktop)
      ...(routes[route].getScreen && !routes[route].skipShim
        ? {
            getScreen: () => {
              if (_cached) {
                return _cached
              }

              _cached = platformWrapper(routes[route].getScreen(), isModal, isLoggedOut)
              return _cached
            },
          }
        : {}),
    }

    return map
  }, {})
}
