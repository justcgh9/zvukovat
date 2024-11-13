import type { AppProps } from "next/app";
import SideMenu from "@/components/SideMenu";
import Player from "@/components/Player";
import { NextThunkDispatch, wrapper } from "@/store";
import { FC, useEffect } from "react";
import './global.scss';
import { checkAuth } from "@/store/action-creators/user";
import { useDispatch } from "react-redux";

const App: FC<AppProps> = ({ Component, pageProps }: AppProps) => {

  const dispatch = useDispatch() as NextThunkDispatch;
  useEffect(() => {
    (async function(){if(localStorage.getItem('token')){
      dispatch(await checkAuth());
    }})()
  }, [])
  
  return (
    <>
      <main id="main">
        <SideMenu />
        <section id="content">
          <Component {...pageProps} />
          <Player/>
        </section>
        
      </main>
      {/* <Container style={{minHeight: '100vh', height: "fit-content", minWidth: '100%', margin: 0, display: 'flex', flexDirection: 'row', padding: 0, overflow:"visible"}} >
        <SideMenu />
        <Container style={{minWidth: '70%', margin: '0 auto', padding: 0}}>
          <Component {...pageProps} />
          <Player/>
        </Container>
        
      
      </Container> */}
      
      </>
  )
}

export default wrapper.withRedux(App);