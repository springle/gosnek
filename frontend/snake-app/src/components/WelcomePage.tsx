import { Button, Callout, Divider, Intent, Spinner } from '@blueprintjs/core';
import React, { Dispatch } from 'react';
import { connect, ConnectedProps, useDispatch } from 'react-redux';
import { RootState } from '..';
import { setConnectionState } from '../redux/actions';
import { ConnectionState } from '../redux/types';

const mapState = (root: RootState) => ({
    connectionState: root.connectionState
});

const connector = connect(mapState);

const onClick = (dispatch: Dispatch<any>) => {
    dispatch(setConnectionState(ConnectionState.LOADING));
    setTimeout(() => {
        dispatch(setConnectionState(ConnectionState.CONNECTED));
    }, 300);
}

const WelcomePage = (props: ConnectedProps<typeof connector>) => {
    const dispatch = useDispatch()
    const spinner = props.connectionState === ConnectionState.LOADING ? <Spinner size={Spinner.SIZE_SMALL}/> : null;
    return (
        <div className='welcome-page bp3-light'>
            <div className='welcome-callout'>
                <Callout intent={Intent.PRIMARY}>
                    <h4 className='bp3-heading'>Welcome</h4>
                    <Divider />
                    <Button text='Join' onClick={() => onClick(dispatch)}></Button>
                    {spinner}
                </Callout>
            </div>
        </div>
    );
}

export default connect(mapState)(WelcomePage);