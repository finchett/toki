import { useEffect, useState } from 'react';
import './App.css';
import { ListRecent } from '../wailsjs/go/main/App';
import { models } from '../wailsjs/go/models';

type Activity = models.Activity & { duration?: string };

function App() {
    const [activities, setActivities] = useState<Activity[]>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);

    const refresh = () => {
        setLoading(true);
        ListRecent(20)
            .then((rows) => {
                setActivities(rows as Activity[]);
                setError(null);
            })
            .catch((err) => setError(String(err)))
            .finally(() => setLoading(false));
    };

    useEffect(refresh, []);

    return (
        <div id="App" style={{ padding: '1.5rem', fontFamily: 'system-ui, sans-serif' }}>
            <header style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'baseline', marginBottom: '1rem' }}>
                <h1 style={{ margin: 0 }}>Toki</h1>
                <button onClick={refresh} disabled={loading}>
                    {loading ? 'Loading…' : 'Refresh'}
                </button>
            </header>

            {error && (
                <div style={{ background: '#5a1f1f', padding: '0.75rem', borderRadius: 4, marginBottom: '1rem' }}>
                    {error}
                </div>
            )}

            {!error && activities.length === 0 && !loading && (
                <p style={{ opacity: 0.7 }}>No activities yet. Run <code>tock start &lt;activity&gt;</code> in your terminal to create one.</p>
            )}

            <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
                {activities.map((a, i) => (
                    <li
                        key={`${a.start_time}-${i}`}
                        style={{
                            padding: '0.75rem 0',
                            borderBottom: '1px solid rgba(255,255,255,0.08)',
                            display: 'grid',
                            gridTemplateColumns: '1fr auto',
                            gap: '0.5rem',
                        }}
                    >
                        <div>
                            <div style={{ fontWeight: 600 }}>{a.description || '(no description)'}</div>
                            <div style={{ opacity: 0.7, fontSize: '0.85rem' }}>
                                {a.project ? `[${a.project}] ` : ''}
                                {new Date(a.start_time).toLocaleString()}
                                {a.end_time ? ` → ${new Date(a.end_time).toLocaleTimeString()}` : ' · running'}
                            </div>
                        </div>
                        <div style={{ fontVariantNumeric: 'tabular-nums', opacity: 0.85 }}>
                            {a.duration ?? ''}
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default App;
