import { Context, MakeStore, createWrapper } from "next-redux-wrapper";
import { Store, createStore } from "redux";
import { RootState, reducer } from "./reducers";

const makeStore: MakeStore<Store<RootState>> = (context: Context) => createStore(reducer);
export const wrapper = createWrapper<Store<RootState>>(makeStore, {debug: true});