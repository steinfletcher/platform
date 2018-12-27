import * as React from 'react'

export type AnyReactComponent<TProps> =
    React.ComponentClass<TProps> | React.FunctionComponent<TProps>

export type WithAnyProps = {
    [key: string]: any;
}

export type WithChildren = {
    children?: any;
}
