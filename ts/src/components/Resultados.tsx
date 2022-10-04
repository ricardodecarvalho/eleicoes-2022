
import { AbrType } from "../types"
import { formatNumber } from "../utils"

const Resultados = (data: AbrType) => {
    return (
        <div className="resultados">
            <p>Última atualização: <strong>{data.dt} {data.ht}</strong></p>
            <p><strong>{data.pst}%</strong> das seções totalizadas</p>
            <table>
                <thead>
                    <tr>
                        <th className='text-center'>Seq</th>
                        <th className='text-center'>Candidato</th>
                        <th className='text-right'>Porcentagem</th>
                        <th className='text-right'>Votos</th>
                        <th className='text-right'>Eleito</th>
                    </tr>
                </thead>
                <tbody>
                    {data.cand && data.cand.map(candidato => (
                        <tr key={candidato.seq}>
                            <td className='text-center'>{candidato.seq}</td>
                            <td className='text-center'>{candidato.n}</td>
                            <td className='text-right'>{candidato.pvap}%</td>
                            <td className='text-right'>{formatNumber(Number(candidato.vap))}</td>
                            <td className='text-center'>{candidato.e}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default Resultados
