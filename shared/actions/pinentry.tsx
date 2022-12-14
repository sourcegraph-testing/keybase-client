import logger from '../logger'
import * as EngineGen from './engine-gen-gen'
import * as PinentryGen from './pinentry-gen'
import * as Constants from '../constants/login'
import * as Container from '../util/container'
import * as RPCTypes from '../constants/types/rpc-gen'
import {getEngine} from '../engine/require'

// stash response while we show the pinentry. The old code kept a map of this but this likely never worked. it seems like
// core sends 0 over and over so it just gets stomped anyways. I have a larger change that removes this kind of flow but
// its not worth implementing now
let _response: EngineGen.Keybase1SecretUiGetPassphrasePayload['payload']['response'] | null = null

const onConnect = async () => {
  try {
    await RPCTypes.delegateUiCtlRegisterSecretUIRpcPromise()
    logger.info('Registered secret ui')
  } catch (error) {
    logger.warn('error in registering secret ui: ', error)
  }
}

const onGetPassword = (_: unknown, action: EngineGen.Keybase1SecretUiGetPassphrasePayload) => {
  logger.info('Asked for password')
  const {response, params} = action.payload
  const {pinentry} = params
  const {prompt, submitLabel, cancelLabel, windowTitle, features, type} = pinentry
  let {retryLabel} = pinentry
  if (retryLabel === Constants.invalidPasswordErrorString) {
    retryLabel = 'Incorrect password.'
  }

  // Stash response
  _response = response

  return PinentryGen.createNewPinentry({
    cancelLabel,
    prompt,
    retryLabel,
    showTyping: features.showTyping,
    submitLabel,
    type,
    windowTitle,
  })
}

const onSubmit = (_: unknown, action: PinentryGen.OnSubmitPayload) => {
  const {password} = action.payload
  if (_response) {
    _response.result({passphrase: password, storeSecret: false})
    _response = null
  }

  return PinentryGen.createClose()
}

const onCancel = () => {
  if (_response) {
    _response.error({code: RPCTypes.StatusCode.scinputcanceled, desc: 'Input canceled'})
    _response = null
  }
  return PinentryGen.createClose()
}

const initPinentry = () => {
  Container.listenAction(PinentryGen.onSubmit, onSubmit)
  Container.listenAction(PinentryGen.onCancel, onCancel)
  getEngine().registerCustomResponse('keybase.1.secretUi.getPassphrase')
  Container.listenAction(EngineGen.keybase1SecretUiGetPassphrase, onGetPassword)
  Container.listenAction(EngineGen.connected, onConnect)
}

export default initPinentry
