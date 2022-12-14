import * as Container from '../../util/container'
import * as Constants from '../../constants/devices'
import * as RouteTreeGen from '../../actions/route-tree-gen'
import DevicePage from '.'

type OwnProps = Container.RouteProps<'devicePage'>

export default Container.connect(
  (state: Container.TypedState, ownProps: OwnProps) => {
    const id = ownProps.route.params?.deviceID ?? ''
    return {iconNumber: Constants.getDeviceIconNumber(state, id), id}
  },
  dispatch => ({
    onBack: () => {
      Container.isMobile && dispatch(RouteTreeGen.createNavigateUp())
    },
  }),
  (stateProps, dispatchProps) => ({
    iconNumber: stateProps.iconNumber,
    id: stateProps.id,
    onBack: dispatchProps.onBack,
  })
)(DevicePage)
