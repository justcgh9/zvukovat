import type { AppProps } from "next/app";
import SideMenu from "@/components/SideMenu";
import Container from "@mui/material/Container";
import Player from "@/components/Player";
import { wrapper } from "@/store";
import { FC } from "react";
import './global.scss';

const App: FC<AppProps> = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Container style={{minHeight: '100%', minWidth: '100%', margin: 0, display: 'flex', flexDirection: 'row', padding: 0}} >
        <SideMenu />
        <Container style={{minWidth: '70%', margin: '0 auto', padding: 0}}>
          <Component {...pageProps} />
        </Container>
        <Player/>
      
      </Container>
      
      </>
  )
}

export default wrapper.withRedux(App);