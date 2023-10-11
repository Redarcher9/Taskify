import { Box } from '@mui/material';
import './App.css';
import Button from '@mui/material/Button';

function App() {

  return (
      <Box sx={{ 
        display: 'flex',
        justifyContent:'center',
        alignItems:'center',
        height:'100vh',
        width:'100vw'}}>

        <Button variant="contained">Hello world</Button>

      </Box>
  );
}

export default App;
