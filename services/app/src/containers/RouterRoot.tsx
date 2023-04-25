import { Suspense } from 'react'
import { useLoaderData, useOutlet, defer, Await } from 'react-router-dom'

export const routerRootLoader = async () => {
  return defer({
    authData: new Promise((resolve) => setTimeout(() => resolve(true), 1000))
  })
}

export const RouterRoot = () => {
  const outlet = useOutlet()
  const { authData } = useLoaderData() as { authData: Promise<any> }

  return (
    <Suspense>
      <Await
        resolve={authData}
        errorElement={<div>Something went wrong!</div>}
        children={(authData) => (
          <div>{outlet}</div>
        )}
      />
    </Suspense>
  )
}
