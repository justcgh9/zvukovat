import type { AppProps } from "next/app";
import Navbar from "@/components/Navbar";
import Container from "@mui/material/Container";
import Player from "@/components/Player";
import { wrapper } from "@/store";
import { FC } from "react";

const App: FC<AppProps> = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Navbar/>
      <Container style={{margin: '90px 0', minHeight: '100%', minWidth: '100%'
          }}>
        <Component {...pageProps} />;      
      </Container>
      <Player/>
    
    </>
  )
}

export default wrapper.withRedux(App);