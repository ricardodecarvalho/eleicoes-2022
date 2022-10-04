type ButtonType = {
    text: string
    disabled: boolean
    onClick: () => void
}
const Button = ({ text, disabled, onClick }: ButtonType) => {
    return (
        <button
            type='button'
            disabled={disabled}
            onClick={onClick}>
            {text}
        </button>
    )
}

export default Button
