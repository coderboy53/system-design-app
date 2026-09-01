import { useCallback, useEffect, useRef, useState } from 'react'
import './App.css'

const BACKEND_ORIGIN = 'http://localhost:9000'

function App() {
  const [status, setStatus] = useState('loading') // 'loading' | 'authenticated' | 'unauthenticated'
  const pollRef = useRef(null)
  const popupRef = useRef(null)

  const checkAuth = useCallback(async () => {
    try {
      const res = await fetch('/auth/status', { credentials: 'include' })
      setStatus(res.status === 200 ? 'authenticated' : 'unauthenticated')
      return res.status === 200
    } catch {
      setStatus('unauthenticated')
      return false
    }
  }, [])

  useEffect(() => {
    checkAuth()
  }, [checkAuth])

  const stopPolling = () => {
    if (pollRef.current) {
      clearInterval(pollRef.current)
      pollRef.current = null
    }
  }

  const handleLogin = () => {
    popupRef.current = window.open(
      `${BACKEND_ORIGIN}/api/login`,
      'google-login',
      'width=500,height=650'
    )

    stopPolling()
    pollRef.current = setInterval(async () => {
      if (popupRef.current && popupRef.current.closed) {
        stopPolling()
      }
      const authed = await checkAuth()
      if (authed) {
        stopPolling()
        if (popupRef.current && !popupRef.current.closed) {
          popupRef.current.close()
        }
      }
    }, 1500)
  }

  const handleLogout = async () => {
    try {
      await fetch('/api/logout', { method: 'POST', credentials: 'include' })
    } catch {
      // backend may respond with an invalid status here; ignore and re-check below
    }
    checkAuth()
  }

  useEffect(() => stopPolling, [])

  if (status === 'loading') {
    return (
      <section id="center">
        <p>Checking session...</p>
      </section>
    )
  }

  if (status === 'unauthenticated') {
    return (
      <section id="center">
        <h1>Welcome</h1>
        <button type="button" className="counter" onClick={handleLogin}>
          Login with Google
        </button>
      </section>
    )
  }

  return (
    <section id="center">
      <h1>Welcome to the app</h1>
      <button type="button" className="counter" onClick={handleLogout}>
        Logout
      </button>
    </section>
  )
}

export default App
