import React from 'react'
import CIcon from '@coreui/icons-react'
import {
  cilBell,
  cilCalculator,
  cilChartPie,
  cilCursor,
  cilDescription,
  cilDrop,
  cilNotes,
  cilPencil,
  cilPuzzle,
  cilSpeedometer,
  cilStar,
} from '@coreui/icons'
import { CNavGroup, CNavItem, CNavTitle } from '@coreui/react'

const _nav = [
  {
    component: CNavItem,
    name: 'Weather',
    to: '/weather',
    icon: <CIcon icon={cilChartPie} customClassName="nav-icon" />,
    badge: {
        color: 'info',
        text: 'NEW',
      },
  },
]

export default _nav
