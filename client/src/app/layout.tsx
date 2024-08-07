import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { AppRouterCacheProvider } from '@mui/material-nextjs/v13-appRouter';
import Navbar from "@/components/Navbar";
import Container from "@mui/material/Container";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "FizzyMusic",
  description: "Some freakingly awesome music player",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <AppRouterCacheProvider>
          <Navbar/>
          <Container style={{margin: '90px 0'}}>
            {children}
          </Container>
        </AppRouterCacheProvider>
        </body>
    </html>
  );
}
