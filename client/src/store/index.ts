import { Context, MakeStore, createWrapper } from "next-redux-wrapper";
import { Action, Store, applyMiddleware, createStore } from "redux";
import { RootState, reducer } from "./reducers";
import { ThunkDispatch, thunk } from "redux-thunk";

const makeStore: MakeStore<Store<RootState>> = (context: Context) => createStore(reducer, applyMiddleware(thunk));
export const wrapper = createWrapper<Store<RootState>>(makeStore, {debug: true});

export type NextThunkDispatch = ThunkDispatch<RootState, void, Action>;