const formatNumber = (number: number, locale = 'pt-BR') => (
    new Intl.NumberFormat(locale).format(number)
)

export default formatNumber