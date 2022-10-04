import Button from './components/Button';
import Resultados from './components/Resultados';

import { useResult } from './hooks'

function App() {
  const {
    resultados,
    loading,
    update,
    error
  } = useResult()

  return (
    <div className="App">

      <Button
        text='Atualizar'
        disabled={loading}
        onClick={update}
      />

      {loading && (
        <p>Buscando novos dados...</p>
      )}

      {!loading && resultados && resultados?.abr && resultados.abr.map((abr, index) => (
        <Resultados key={index} {...abr} />
      ))}

      {!loading && error && error?.message && (
        <div className="error">{error.message}</div>
      )}

    </div>
  );
}

export default App;
