// store.js
import { Action, configureStore, ThunkDispatch } from '@reduxjs/toolkit';
import { createWrapper } from 'next-redux-wrapper';
import {rootReducer, RootState } from './reducers/index';

const makeStore = () => configureStore({
  reducer: rootReducer,
});

export const wrapper = createWrapper(makeStore);


export type NextThunkDispatch = ThunkDispatch<RootState, void, Action>;