import * as React from 'react'
import type {Position} from './relative-popup-hoc.types'
import type {StylesCrossPlatform} from '../styles'

export type Props = {
  children: React.ReactNode
  className?: string
  containerStyle?: StylesCrossPlatform
  visible: boolean
  attachTo?: () => any
  // applies on desktop only. Mobile is always centered in the screen
  position?: Position
}

export default class extends React.Component<Props> {}
