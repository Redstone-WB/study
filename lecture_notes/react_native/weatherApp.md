### WeatherApp (Nodad coders)

1. App에서 default Flex direction은 column임. (세로방향)
cf. Web은 row (가로방향)

StyleSheet.create({
    container :{
        flex: 1, 
        flexDirection : "row", ## 이렇게 바꿀 수 있음.
    }
})


2. 'flex :1' 의미 : 모든 공간을 사용할 수 있음.
'flex : 2' 인 yellowview와 'flew : 1'인 blueview가 있다면,
yellowview 가 화면에서 2/3 을 차지하게 됨.
-> width / height 사용하는 것보다 위와 같이 flex를 사용하면 비율로 조정이 가능함. 이는 크기가 서로 다른 여러 모바일 기기 호환에 유리.
