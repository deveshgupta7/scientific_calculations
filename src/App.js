import React, { useState } from 'react';
import axios from 'axios';

const Calculator = () => {
  const [display, setDisplay] = useState('');
  const [operation, setOperation] = useState('');
  const [firstNumber, setFirstNumber] = useState(null);

  const handleNumberClick = (num) => {
    setDisplay(display + num);
  };

  const handleOperationClick = (op) => {
    setOperation(op);
    setFirstNumber(parseFloat(display));
    setDisplay('');
  };

  const handleCalculate = async () => {
    const secondNumber = parseFloat(display);
    
    try {
      const response = await axios.post('http://localhost:8080/calculate', {
        operation: operation,
        numbers: [firstNumber, secondNumber],
      });

      setDisplay(response.data.result.toString());
    } catch (error) {
      setDisplay('Error');
    }
  };

  const handleClear = () => {
    setDisplay('');
    setOperation('');
    setFirstNumber(null);
  };

  return (
    <div className="calculator">
      <input type="text" value={display} readOnly />
      <div className="keypad">
        {[7, 8, 9, 4, 5, 6, 1, 2, 3, 0].map((num) => (
          <button key={num} onClick={() => handleNumberClick(num.toString())}>
            {num}
          </button>
        ))}
        <button onClick={() => handleOperationClick('add')}>+</button>
        <button onClick={() => handleOperationClick('subtract')}>-</button>
        <button onClick={() => handleOperationClick('multiply')}>*</button>
        <button onClick={() => handleOperationClick('divide')}>/</button>
        <button onClick={() => handleOperationClick('power')}>^</button>
        <button onClick={() => handleOperationClick('sqrt')}>√</button>
        <button onClick={() => handleOperationClick('sin')}>sin</button>
        <button onClick={() => handleOperationClick('cos')}>cos</button>
        <button onClick={() => handleOperationClick('tan')}>tan</button>
        <button onClick={handleCalculate}>=</button>
        <button onClick={handleClear}>C</button>
      </div>
    </div>
  );
};

export default Calculator;
