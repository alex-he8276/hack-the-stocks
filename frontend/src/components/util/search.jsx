import * as React from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { stocks } from '../../constants/sandp500'
const Search = ({value, inputValue, setValue, setInputValue}) => {
    return (
        <Autocomplete
          value={value}
          disablePortal
          id="combo-box-demo"
          options={stocks}
          sx={{ width: 300 }}
          renderInput={(params) => <TextField {...params} label="Company" />}
          onChange={(event, newValue) => {
            console.log(newValue)
            setValue(newValue.Symbol);
          }}
        />
      );
};

export default Search;
