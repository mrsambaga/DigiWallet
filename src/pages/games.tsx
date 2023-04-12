import React, { useEffect, useState } from 'react';
import Title from '../components/title';
import useFetchGet from '../hooks/useFetchGet';
import useFetchPost from '../hooks/useFetchPost';
import { notifyError } from '../components/notification';
import { ProfileResponse } from '../types/types';
import { GetCookie } from '../helper/cookies';
import '../styles/games/games.css';
import { NavLink } from 'react-router-dom';
import { shuffle } from '../helper/shuffle';

const Games: React.FC = () => {
  const token = GetCookie('token');

  const arrayNum: number[] = [1, 2, 3, 4, 5, 6, 7, 8, 9];
  const [shuffledArr, setShuffledArr] = useState(shuffle(arrayNum));
  const [prize, setPrize] = useState(shuffledArr);

  const [submit, setSubmit] = useState(false);
  const [selectedId, setSelectedId] = useState(0);
  const [boxClass, setBoxClass] = useState('box-inactive');
  const boxOnClickHandler = (id: number) => {
    if (chance > 0 && boxClass == 'box-inactive') {
      setBoxClass('box-active');
      setSelectedId(id);
      setSubmit(true);
      console.log('BOX ID :', id);
      return;
    } else if (chance == 0 && boxClass == 'box-inactive') {
      setBoxClass('box-no-chance');
      return;
    }

    setShuffledArr(shuffle(shuffledArr));
    setBoxClass('box-inactive');
  };

  const { out: gamesOut, error: gamesErr } = useFetchPost(
    'http://localhost:8000/games',
    { box_id: selectedId },
    submit,
    () => setSubmit(false),
    token!,
  );

  useEffect(() => {
    if (gamesErr != null) {
      notifyError(gamesErr.response?.data?.message || gamesErr.message);
    } else if (gamesOut != null) {
      console.log(gamesOut.data);
      console.log('SHUFFLE', shuffledArr);
      const updatedPrizeArr = shuffledArr.map((id) => {
        const box = gamesOut.data.find((b: any) => b.box_id === id);
        return box != null ? box.prize : null;
      });
      setPrize(updatedPrizeArr);
      console.log('PRIZE', updatedPrizeArr);
    }
  }, [gamesOut, gamesErr]);

  const [profileResponse, setProfileResponse] = useState<ProfileResponse>({
    Balance: 0,
    Email: '',
    UserId: 0,
    UserName: '',
    WalletNumber: 0,
  });
  const { out: profileOut, error: profileErr } = useFetchGet(
    `http://localhost:8000/profile`,
    token!,
    submit,
  );

  useEffect(() => {
    if (profileErr) {
      notifyError(profileErr.response?.data?.message || profileErr.message);
      return;
    }

    if (profileOut != null) {
      const profileResponse: ProfileResponse = {
        Balance: profileOut.data.balance,
        Email: profileOut.data.email,
        UserId: profileOut.data.user_id,
        UserName: profileOut.data.user_name,
        WalletNumber: profileOut.data.wallet_number,
      };

      localStorage.setItem('wallet_number', profileOut.data.wallet_number);
      setProfileResponse(profileResponse);
    }

    console.log('INI EKSEKUSI');
  }, [profileOut, profileErr]);

  const [chance, setChance] = useState(0);
  const { out: chanceOut, error: chanceErr } = useFetchGet(
    `http://localhost:8000/chance`,
    token!,
    submit,
  );

  useEffect(() => {
    if (chanceErr) {
      notifyError(chanceErr.response?.data?.message || chanceErr.message);
      return;
    }

    if (chanceOut != null) {
      setChance(chanceOut.data.Chance);
    }
  }, [chanceOut, chanceErr]);

  return (
    <div className="games">
      <div className="games__container">
        <Title idName="games-title" contentProps={profileResponse!} />
        <div className="games__container__table">
          <h1>Games</h1>
          <p>Choose random box below to get reward!</p>
          <p>Chance : {chance}</p>
          <NavLink to="/leaderboard">Check Leaderboard</NavLink>
          <div className="games__container__table__box">
            <button
              onClick={() => boxOnClickHandler(shuffledArr[0])}
              className={boxClass}
              id={
                selectedId == shuffledArr[0] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[0] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[1])}
              className={boxClass}
              id={
                selectedId == shuffledArr[1] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[1] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[2])}
              className={boxClass}
              id={
                selectedId == shuffledArr[2] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[2] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[3])}
              className={boxClass}
              id={
                selectedId == shuffledArr[3] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[3] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[4])}
              className={boxClass}
              id={
                selectedId == shuffledArr[4] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[4] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[5])}
              className={boxClass}
              id={
                selectedId == shuffledArr[5] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[5] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[6])}
              className={boxClass}
              id={
                selectedId == shuffledArr[6] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[6] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[7])}
              className={boxClass}
              id={
                selectedId == shuffledArr[7] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[7] : ''}
            </button>
            <button
              onClick={() => boxOnClickHandler(shuffledArr[8])}
              className={boxClass}
              id={
                selectedId == shuffledArr[8] && boxClass == 'box-active'
                  ? 'selectedBox'
                  : ''
              }
            >
              {boxClass == 'box-active' ? prize[8] : ''}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Games;
