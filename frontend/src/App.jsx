import { useState, useEffect } from 'react';
import {
  Container,
  Paper,
  Box,
  Typography,
  IconButton,
  ButtonGroup,
  Button,
  Chip,
  Stack,
} from '@mui/material';
import {
  PlayArrow,
  Pause,
  VolumeUp,
  VolumeOff,
  GraphicEq,
} from '@mui/icons-material';

function App() {
  const [isPaused, setIsPaused] = useState(false);
  const [playbackPosition, setPlaybackPosition] = useState('0s');
  const [currentSpeed, setCurrentSpeed] = useState(1);
  const [isMuted, setIsMuted] = useState(false);
  const [isBoostActive, setIsBoostActive] = useState(false);

  const speedOptions = [
    { label: '0.5x', value: 0.5, key: '0' },
    { label: '1x', value: 1, key: '1' },
    { label: '2x', value: 2, key: '2' },
    { label: '3x', value: 3, key: '3' },
    { label: '4x', value: 4, key: '4' },
  ];

  // Keyboard shortcuts
  useEffect(() => {
    const handleKeyPress = (e) => {
      switch (e.key) {
        case 'Enter':
          handlePlayPause();
          break;
        case '0':
        case '1':
        case '2':
        case '3':
        case '4':
          const speedOption = speedOptions.find(s => s.key === e.key);
          if (speedOption) handleSpeedChange(speedOption.value);
          break;
        case 'd':
        case 'D':
          handleBoostToggle();
          break;
        case 'm':
        case 'M':
          handleMuteToggle();
          break;
      }
    };

    window.addEventListener('keydown', handleKeyPress);
    return () => window.removeEventListener('keydown', handleKeyPress);
  }, []);

  const handlePlayPause = () => {
    setIsPaused(prev => !prev);
    // TODO: Send command to backend
    console.log('Play/Pause toggled');
  };

  const handleSpeedChange = (speed) => {
    setCurrentSpeed(speed);
    // TODO: Send command to backend
    console.log('Speed changed to:', speed);
  };

  const handleMuteToggle = () => {
    setIsMuted(prev => !prev);
    // TODO: Send command to backend
    console.log('Mute toggled');
  };

  const handleBoostToggle = () => {
    setIsBoostActive(prev => !prev);
    // TODO: Send command to backend
    console.log('Boost toggled');
  };

  return (
    <Container maxWidth="sm" sx={{ mt: 8 }}>
      <Paper elevation={6} sx={{ p: 4, borderRadius: 3 }}>
        <Typography variant="h4" component="h1" gutterBottom align="center" fontWeight="bold">
          🎵 Smart Audio Player
        </Typography>

        {/* Position Display */}
        <Box sx={{ my: 4, textAlign: 'center' }}>
          <Typography variant="h2" component="div" color="primary" fontWeight="bold">
            {playbackPosition}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Playback Position
          </Typography>
        </Box>

        {/* Play/Pause Button */}
        <Box sx={{ display: 'flex', justifyContent: 'center', my: 3 }}>
          <IconButton
            onClick={handlePlayPause}
            sx={{
              bgcolor: 'primary.main',
              color: 'white',
              width: 80,
              height: 80,
              '&:hover': { bgcolor: 'primary.dark' },
            }}
          >
            {isPaused ? <PlayArrow sx={{ fontSize: 50 }} /> : <Pause sx={{ fontSize: 50 }} />}
          </IconButton>
        </Box>

        {/* Speed Control */}
        <Box sx={{ my: 3 }}>
          <Typography variant="subtitle1" gutterBottom align="center" fontWeight="medium">
            Playback Speed
          </Typography>
          <ButtonGroup fullWidth variant="contained" size="large">
            {speedOptions.map((option) => (
              <Button
                key={option.key}
                onClick={() => handleSpeedChange(option.value)}
                variant={currentSpeed === option.value ? 'contained' : 'outlined'}
                color={currentSpeed === option.value ? 'primary' : 'inherit'}
              >
                {option.label}
              </Button>
            ))}
          </ButtonGroup>
        </Box>

        {/* Audio Controls */}
        <Stack direction="row" spacing={2} sx={{ mt: 4 }} justifyContent="center">
          <Chip
            icon={isMuted ? <VolumeOff /> : <VolumeUp />}
            label={isMuted ? 'Unmute (M)' : 'Mute (M)'}
            onClick={handleMuteToggle}
            color={isMuted ? 'error' : 'default'}
            variant={isMuted ? 'filled' : 'outlined'}
            sx={{ px: 2, py: 3 }}
          />
          <Chip
            icon={<GraphicEq />}
            label={isBoostActive ? 'Boost ON (D)' : 'Boost OFF (D)'}
            onClick={handleBoostToggle}
            color={isBoostActive ? 'success' : 'default'}
            variant={isBoostActive ? 'filled' : 'outlined'}
            sx={{ px: 2, py: 3 }}
          />
        </Stack>

        {/* Keyboard Shortcuts */}
        <Box sx={{ mt: 4, p: 2, bgcolor: 'grey.100', borderRadius: 2 }}>
          <Typography variant="caption" display="block" gutterBottom fontWeight="medium">
            Keyboard Shortcuts:
          </Typography>
          <Typography variant="caption" display="block" color="text.secondary">
            <strong>Enter:</strong> Play/Pause • <strong>0-4:</strong> Speed • <strong>M:</strong> Mute •{' '}
            <strong>D:</strong> Boost
          </Typography>
        </Box>
      </Paper>
    </Container>
  );
}

export default App;
