import React from 'react'

const Charts = React.lazy(() => import('./views/charts/Charts'))


const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/weather', name: 'Charts', element: Charts },
]

export default routes
