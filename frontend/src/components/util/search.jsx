import * as React from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import { stocks } from '../../constants/sandp500'
const Search = ({setValue, setInputValue}) => {
    return (
        <Autocomplete
          disablePortal
          id="combo-box-demo"
          options={stocks}
          sx={{ width: 300 }}
          renderInput={(params) => <TextField {...params} label="Company" />}
          onChange={(event, newValue) => {
            console.log(newValue)
            setValue(newValue.Symbol);
            setInputValue(newValue.label)
          }}
        />
      );
};

export default Search;
