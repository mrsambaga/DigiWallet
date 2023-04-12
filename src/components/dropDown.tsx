import React from 'react';
import '../styles/dropdown/dropdown.css';
import { DropdownOption } from '../types/types';

interface DropdownProps {
  onChange: (value: React.ChangeEvent<HTMLSelectElement>) => void;
  label: string;
  dropdownOptions: DropdownOption[];
}

const Dropdown: React.FC<DropdownProps> = ({
  onChange,
  label,
  dropdownOptions,
}) => {
  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    onChange(event);
  };

  return (
    <div className="dropdown">
      <label className="dropdown__label">{label}</label>
      <select
        onChange={handleSelectChange}
        className={dropdownOptions[0].value}
        defaultValue={dropdownOptions[0].value}
      >
        <option disabled value="" className="option"></option>
        {dropdownOptions.map((option) => (
          <option key={option.value} value={option.value}>
            {option.content}
          </option>
        ))}
      </select>
    </div>
  );
};

export default Dropdown;
