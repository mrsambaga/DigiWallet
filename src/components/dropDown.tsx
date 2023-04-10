import React from 'react';
import '../styles/dropdown/dropdown.css';

interface DropdownProps {
  onChange: (value: React.ChangeEvent<HTMLSelectElement>) => void;
  label: string;
}

const Dropdown: React.FC<DropdownProps> = ({ onChange, label }) => {
  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    onChange(event);
  };

  return (
    <div className="dropdown">
      <label className="dropdown__label">{label}</label>
      <select
        onChange={handleSelectChange}
        className="dropdown__select"
        defaultValue={'bank transfer'}
      >
        <option disabled value="" className="option"></option>
        <option value="bank transfer">Bank Transfer</option>
        <option value="credit card">Credit Card</option>
        <option value="cash">Cash</option>
      </select>
    </div>
  );
};

export default Dropdown;
