import { useState, useEffect, useCallback } from 'react'
import axios, { AxiosError } from 'axios'

import { ResultadosType } from '../types'

const URL = "https://resultados.tse.jus.br/oficial/ele2022/544/dados/br/br-c0001-e000544-v.json"

const useResult = () => {
  const [resultados, setResultados] = useState<ResultadosType|null>(null)
  const [error, setError] = useState<AxiosError|null>()
  const [loading, setLoading] = useState<boolean>(false)

  const update = useCallback(() => {
    setLoading(true)
    axios.get(URL)
    .then(function (response) {
      setLoading(false)
      setResultados(response.data);
      setError(null);
    })
    .catch(function (error) {
      console.log(error)
      setLoading(false)
      setResultados(null)
      setError(error)
    });
  }, [])

  useEffect(() => {
    update()
  }, [update])

  return { resultados, error, update, loading }
}

export default useResult
