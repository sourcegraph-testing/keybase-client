import * as React from 'react'
import {StylesCrossPlatform} from '../styles'

export type Props = {
  children?: React.ReactNode
  style?: StylesCrossPlatform
}

export default class SafeAreaView extends React.Component<Props> {}
export class SafeAreaViewTop extends React.Component<Props> {}
export declare function useSafeAreaInsets(): {
  bottom: number
  left: number
  right: number
  top: number
}
