interface Props {
  tip?: string
  size?: 'small' | 'default' | 'large'
  fullScreen?: boolean
}

const sizeMap = {
  small: '20px',
  default: '32px',
  large: '48px',
}

export default function Loading({ tip = '加载中...', size = 'default', fullScreen = false }: Props): JSX.Element {
  const spinnerSize = sizeMap[size]

  const containerStyle: React.CSSProperties = fullScreen
    ? {
        position: 'fixed',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        background: 'rgba(255, 255, 255, 0.9)',
        zIndex: 9999,
      }
    : {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        padding: '2rem',
      }

  return (
    <div style={containerStyle}>
      <div
        style={{
          width: spinnerSize,
          height: spinnerSize,
          border: '3px solid var(--border)',
          borderTopColor: 'var(--primary)',
          borderRadius: '50%',
          animation: 'spin 0.8s linear infinite',
        }}
      />
      {tip && (
        <span style={{ marginTop: '0.75rem', color: 'var(--text-muted)', fontSize: '0.875rem' }}>
          {tip}
        </span>
      )}
      <style>{`
        @keyframes spin {
          to { transform: rotate(360deg); }
        }
      `}</style>
    </div>
  )
}