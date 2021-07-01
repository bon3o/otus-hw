package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var textWithCommas = `Этот текст,написанный без пробелов после 
	запятых,должен без ошибок пройти частотный анализ.Также,ТекСт 
	написан словамис рандомным регистром,что должно показать,что 
	это не 	имеет значения. Текст,теКст,ТЕКСТ,текст!!!`

var textLorem = `Maecenas ipsum velit, consectetuer eu, lobortis ut, 
	dictum at, dui. In rutrum. Sed ac dolor sit amet purus malesuada 
	congue. In laoreet, magna id viverra tincidunt, sem odio bibendum 
	justo, vel imperdiet sapien wisi sed libero. Suspendisse sagittis 
	ultrices augue. Mauris metus. Nunc dapibus tortor vel mi dapibus 
	sollicitudin. Etiam posuere lacus quis dolor. Praesent id justo 
	in neque elementum ultrices. Class aptent taciti sociosqu ad litora 
	torquent per conubia nostra, per inceptos hymenaeos. In convallis. 
	Fusce suscipit libero eget elit. Praesent vitae arcu tempor neque 
	lacinia pretium. Morbi imperdiet, mauris ac auctor dictum, nisl 
	ligula egestas nulla, et sollicitudin sem purus in lacus. 
	Morbi leo mi, nonummy eget, tristique non, rhoncus non, leo. 
	Nullam faucibus mi quis velit. Integer in sapien. Fusce tellus odio, 
	dapibus id, fermentum quis, suscipit id, erat. Fusce aliquam vestibulum 
	ipsum. Aliquam erat volutpat. Pellentesque sapien. Cras elementum. 
	Nulla pulvinar eleifend sem. Cum sociis natoque penatibus et magnis 
	dis parturient montes, nascetur ridiculus mus. Quisque porta. 
	Vivamus porttitor turpis ac leo. Aenean placerat. In vulputate 
	urna eu arcu. Aliquam erat volutpat. Suspendisse potenti. Morbi 
	mattis felis at nunc. Duis viverra diam non justo. In nisl. Nullam 
	sit amet magna in magna gravida vehicula. Mauris tincidunt sem sed 
	arcu. Nunc posuere. Nullam lectus justo, vulputate eget, mollis sed, 
	tempor sed, magna. Cum sociis natoque penatibus et magnis dis 
	parturient montes, nascetur ridiculus mus. Etiam neque. Curabitur 
	ligula sapien, pulvinar a, vestibulum quis, facilisis vel, sapien. 
	Nullam eget nisl. Donec vitae arcu. Lorem ipsum dolor sit amet, 
	consectetuer adipiscing elit. Morbi gravida libero nec velit. 
	Morbi scelerisque luctus velit. Etiam dui sem, fermentum vitae, 
	sagittis id, malesuada in, quam. Proin mattis lacinia justo. 
	Vestibulum facilisis auctor urna. Aliquam in lorem sit amet leo 
	accumsan lacinia. Integer rutrum, orci vestibulum ullamcorper ultricies, 
	lacus quam ultricies odio, vitae placerat pede sem sit amet enim. 
	Phasellus et lorem id felis nonummy placerat. Fusce dui leo, imperdiet in, 
	aliquam sit amet, feugiat eu, orci. Aenean vel massa quis mauris 
	vehicula lacinia. Quisque tincidunt scelerisque libero. Maecenas libero. 
	Etiam dictum tincidunt diam. Donec ipsum massa, ullamcorper in, auctor et, 
	scelerisque sed, est. Suspendisse nisl. Sed convallis magna eu sem. 
	Cras pede libero, dapibus nec, pretium sit amet, tempor quis, urna.
	In sem justo, commodo ut, suscipit at, pharetra vitae, orci. Duis 
	sapien nunc, commodo et, interdum suscipit, sollicitudin et, dolor. 
	Pellentesque habitant morbi tristique senectus et netus et malesuada 
	fames ac turpis egestas. Aliquam id dolor. Class aptent taciti 
	sociosqu ad litora torquent per conubia nostra, per inceptos hymenaeos. 
	Mauris dictum facilisis`

var shortText = "Очень, очень короткий рассказ..."

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
	t.Run("different register test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"текст",    // 6
				"без",      // 2
				"что",      // 2
				"анализ",   // 1
				"должен",   // 1
				"должно",   // 1
				"запятых",  // 1
				"значения", // 1
				"имеет",    // 1
				"написан",  // 1
			}
			require.Equal(t, expected, Top10(textWithCommas))
		}
	})
	t.Run("different register test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"in",      // 14
				"et",      // 9
				"sem",     // 8
				"amet",    // 7
				"id",      // 7
				"sed",     // 7
				"sit",     // 7
				"aliquam", // 6
				"justo",   // 6
				"libero",  // 6
			}
			require.Equal(t, expected, Top10(textLorem))
		}
	})
	t.Run("different register test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"очень",    // 2
				"короткий", // 1
				"рассказ",  // 1
			}
			require.Equal(t, expected, Top10(shortText))
		}
	})
}
