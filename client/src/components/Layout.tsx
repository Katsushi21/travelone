import React from "react"
import Head from "next/head"
import { AppBar, Toolbar, Typography, Container } from "@mui/material"

export const Layout = ({ children }) => {
  return (
    <>
      <div>
        <Head></Head>
        <AppBar>
          <Toolbar></Toolbar>
        </AppBar>
        <Container>{children}</Container>
      </div>
    </>
  )
}
