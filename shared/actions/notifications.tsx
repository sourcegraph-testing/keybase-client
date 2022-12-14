import * as Constants from '../constants/notifications'
import * as ConfigGen from './config-gen'
import * as EngineGen from './engine-gen-gen'
import * as NotificationsGen from './notifications-gen'
import * as FsGen from './fs-gen'
import * as RPCTypes from '../constants/types/rpc-gen'
import * as Container from '../util/container'
import logger from '../logger'
import {isMobile} from '../constants/platform'

const setupNotifications = async () => {
  try {
    await RPCTypes.notifyCtlSetNotificationsRpcPromise({
      channels: {
        allowChatNotifySkips: true,
        app: true,
        audit: true,
        badges: true,
        chat: true,
        chatattachments: true,
        chatdev: false,
        chatemoji: false,
        chatemojicross: false,
        chatkbfsedits: false,
        deviceclone: false,
        ephemeral: false,
        favorites: false,
        featuredBots: true,
        kbfs: true,
        kbfsdesktop: !isMobile,
        kbfslegacy: false,
        kbfsrequest: false,
        kbfssubscription: true,
        keyfamily: false,
        paperkeys: false,
        pgp: true,
        reachability: true,
        runtimestats: true,
        saltpack: true,
        service: true,
        session: true,
        team: true,
        teambot: false,
        tracking: true,
        users: true,
        wallet: true,
      },
    })
  } catch (error) {
    if (error != null) {
      logger.warn('error in toggling notifications: ', error)
    }
  }
}

const createBadgeState = (_: unknown, action: EngineGen.Keybase1NotifyBadgesBadgeStatePayload) =>
  NotificationsGen.createReceivedBadgeState({badgeState: action.payload.params.badgeState})

const receivedBadgeState = (
  state: Container.TypedState,
  action: NotificationsGen.ReceivedBadgeStatePayload
) => {
  const counts = Constants.badgeStateToBadgeCounts(state, action.payload.badgeState)
  return [
    counts && NotificationsGen.createSetBadgeCounts({counts}),
    !isMobile && Constants.shouldTriggerTlfLoad(action.payload.badgeState) && FsGen.createFavoritesLoad(),
  ]
}

const receivedRootAuditError = (_: unknown, action: EngineGen.Keybase1NotifyAuditRootAuditErrorPayload) =>
  ConfigGen.createGlobalError({
    globalError: new Error(`Keybase is buggy, please report this: ${action.payload.params.message}`),
  })

const receivedBoxAuditError = (_: unknown, action: EngineGen.Keybase1NotifyAuditBoxAuditErrorPayload) =>
  ConfigGen.createGlobalError({
    globalError: new Error(
      `Keybase had a problem loading a team, please report this with \`keybase log send\`: ${action.payload.params.message}`
    ),
  })

const initNotifications = () => {
  Container.listenAction(NotificationsGen.receivedBadgeState, receivedBadgeState)
  Container.listenAction(EngineGen.keybase1NotifyAuditRootAuditError, receivedRootAuditError)
  Container.listenAction(EngineGen.keybase1NotifyAuditBoxAuditError, receivedBoxAuditError)
  Container.listenAction(EngineGen.connected, setupNotifications)
  Container.listenAction(EngineGen.keybase1NotifyBadgesBadgeState, createBadgeState)
}

export default initNotifications
