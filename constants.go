package main

var NUMBER_OF_USERS = 20

var RAW_TITLES = "resources/rawTitles.txt"
var RAW_PARAGRAPHS = "resources/rawParagraphs.txt"

var LINE_BREAK = "\n--#######################################################################################\n"
var NEW_LINE = "\n"

var POST_TITLE []string = []string{}

var RESPONSES []string = []string{
	`Lorem ipsum dolor sit amet, consectetur adipiscing elit.`,
	`Nunc in orci maximus, vestibulum eros a, elementum dui.`,
	`Aliquam non tortor vitae purus tempus sollicitudin et ut quam.`,
	`Aliquam volutpat nisi at urna iaculis tristique.`,
	`Quisque vitae enim et purus lacinia congue.`,
	`Donec consequat magna at faucibus lacinia.`,
	`Nam porttitor ex ac venenatis efficitur.`,
	`Proin euismod lacus sit amet aliquam maximus.`,
	`Sed vel dolor quis dui sollicitudin scelerisque.`,
	`Pellentesque ac nisl molestie, congue ipsum quis, lobortis nunc.`,
	`Sed in sem sit amet tellus venenatis dignissim.`,
	`Praesent rhoncus libero sed ante mattis fermentum.`,
	`Nunc ut augue sed tellus vulputate pretium quis eget sem.`,
	`Integer vitae diam et enim ultrices consectetur.`,
	`Phasellus commodo enim vel feugiat porta.`,
	`Sed pellentesque orci a volutpat interdum.`,
	`Etiam sollicitudin diam a quam ullamcorper lobortis.`,
	`Sed ornare est ut lectus blandit viverra.`,
	`Praesent placerat neque ut sapien laoreet tincidunt.`,
	`Nullam tincidunt libero at sapien iaculis, vestibulum sagittis dolor elementum.`,
	`Pellentesque consequat dolor vitae vehicula malesuada.`,
	`Fusce in mi sit amet turpis hendrerit laoreet.`,
	`Phasellus nec sapien lacinia, lacinia erat id, condimentum odio.`,
	`Praesent congue ante non luctus euismod.`,
	`Integer dictum ligula sit amet ullamcorper aliquet.`,
	`Etiam ullamcorper orci quis dignissim laoreet.`,
	`Vestibulum rhoncus nunc eu dui dignissim elementum.`,
	`Morbi at augue ornare, bibendum nunc eu, mattis sapien.`,
	`Nam vulputate arcu non quam gravida, non consequat justo rutrum.`,
	`Nam laoreet massa ac justo bibendum vehicula.`,
	`Morbi eget ex vel elit volutpat rhoncus quis sed orci.`,
	`In at ante in mauris maximus iaculis.`,
	`Aliquam eget orci laoreet, rutrum turpis pharetra, semper sapien.`,
	`Quisque egestas eros in rutrum tincidunt.`,
	`Proin congue sapien non erat dapibus suscipit.`,
	`Nullam cursus neque sed nisi fringilla, sed porttitor nulla dignissim.`,
	`Integer facilisis justo sit amet neque accumsan, sit amet tincidunt odio tristique.`,
	`Mauris id erat sed eros mattis rhoncus.`,
	`Pellentesque quis sapien in ipsum efficitur bibendum non elementum turpis.`,
	`Nunc sit amet tortor ac lacus lacinia aliquam quis eu lectus.`,
	`Pellentesque lobortis tortor vitae mi suscipit, eu semper ante sagittis.`,
	`Nullam lacinia diam eget dolor sodales, eget fringilla libero pretium.`,
	`Suspendisse et magna lobortis, bibendum nibh vel, porttitor ante.`,
	`Integer rutrum nisi sed vestibulum dapibus.`,
	`Nunc porta urna sit amet nisl pulvinar efficitur.`,
	`Mauris dictum lacus vitae felis molestie ullamcorper.`,
	`Curabitur luctus turpis a nulla congue, a dignissim erat sodales.`,
	`Nunc condimentum est vitae maximus tristique.`,
	`Praesent tempor ex eget hendrerit ultricies.`,
	`Quisque auctor elit quis mi aliquet, vel sodales ipsum mattis.`,
	`Suspendisse pulvinar augue nec sem auctor venenatis.`,
	`Integer gravida mauris sit amet ante efficitur, ut varius sapien euismod.`,
	`Maecenas aliquet enim non sem pharetra elementum.`,
	`Integer mollis nisi quis ligula fermentum, quis posuere dolor aliquam.`,
	`Etiam facilisis quam eget odio pharetra pulvinar.`,
	`Cras vitae justo elementum ex porttitor vestibulum eu vel ante.`,
	`Vestibulum id sem at nunc commodo ultrices.`,
	`Nullam malesuada nisi nec nibh tempor sodales.`,
	`Sed tristique quam maximus, placerat lacus eget, blandit mi.`,
	`Nunc blandit mi non neque eleifend, semper fermentum tortor consequat.`,
	`Etiam finibus ex quis nibh egestas pellentesque.`,
	`Sed ultrices magna id risus efficitur, nec consequat massa convallis.`,
	`Ut a odio at risus pharetra viverra ac id ex.`,
	`Quisque malesuada orci eget ligula consectetur tempor.`,
	`Vivamus vel nulla in risus congue scelerisque.`,
	`Nam molestie lectus ut faucibus euismod.`,
	`Nulla sollicitudin massa id ante molestie, a tempus sapien cursus.`,
	`Aenean elementum massa eu libero molestie sagittis.`,
	`Cras cursus libero dignissim dolor dictum rhoncus.`,
	`Nullam porta velit ac erat rhoncus, nec mollis enim dictum.`,
	`Nulla pellentesque est ac odio accumsan commodo.`,
	`Nunc pellentesque mauris eu ornare condimentum.`,
	`Mauris at justo cursus, consequat leo porta, scelerisque odio.`,
	`Maecenas posuere ex sed egestas luctus.`,
	`Donec vel tellus a est sodales vulputate.`,
}

var TEMP []string = []string{
	`Varius laoreet arcu sociosqu, nullam tellus metus nulla at mi etiam hymenaeos enim sapien.`,
	`Varius elementum sapien non vehicula feugiat id elit rhoncus blandit posuere lacinia turpis lectus.`,
	`Nisl aliquet urna justo convallis leo pretium pulvinar sem neque dolor nisi.`,
	`Mi mattis posuere molestie nulla. Habitasse risus primis iaculis nisi.`,
	`Vulputate auctor ante Augue auctor semper commodo quam sapien hendrerit bibendum ad tempus eu.`,
	`Gravida taciti nam imperdiet blandit magnis ultrices gravida rutrum metus.`,
	`Pulvinar maecenas euismod. Ac. Molestie porta. Pede semper ligula elit semper proin magna, dolor.`,
	`Hendrerit sociosqu magnis parturient suspendisse integer ultrices sociis vulputate varius.`,
	`Posuere augue placerat molestie nibh taciti.`,
	`Tempor imperdiet varius fermentum risus platea.`,
	`Iaculis faucibus, pede elementum. Sollicitudin torquent duis habitant.`,
	`Tristique imperdiet at nibh lorem habitasse pede sodales.`,
	`A primis volutpat ligula dui integer penatibus.`,
	`Dis sapien. Torquent egestas justo, inceptos. Praesent.`,
	`Elit quisque pede litora fames. Sed inceptos elementum, interdum conubia nostra lectus congue.`,
	`Sem justo conubia potenti per dui cras. Sodales velit, parturient nullam porta condimentum.`,
	`Quisque placerat odio sollicitudin in integer, nec tempor nisi.`,
	`Odio tincidunt. Platea. Potenti praesent quisque. Semper maecenas egestas augue aptent.`,
	`Hendrerit porta. Molestie. Nisi platea maecenas.`,
	`Nascetur mollis euismod porttitor aenean. Nisl litora mollis venenatis habitasse mattis ornare porta.`,
	`Feugiat consectetuer cras. Egestas eleifend tincidunt bibendum non integer Dui, imperdiet parturient risus.`,
	`Netus hendrerit varius magnis, pellentesque. Fusce urna.`,
	`Commodo nonummy et turpis feugiat vel vehicula viverra id magna a viverra.`,
	`Amet sed Parturient tortor dapibus augue.`,
	`Sociis condimentum sed viverra donec purus. Condimentum ac molestie mi sem consectetuer cras at.`,
	`Non etiam sollicitudin maecenas senectus platea.`,
	`Fusce Auctor facilisi volutpat. Eu sollicitudin. Porttitor Hendrerit.`,
	`Scelerisque eros euismod eget dolor orci porta platea cum proin arcu rutrum felis.`,
	`Ad dolor ut semper orci adipiscing orci.`,
	`Posuere aptent lacus purus nisl pharetra interdum. Duis. Interdum nullam blandit cum ad.`,
	`Mus mauris. Ut luctus. Accumsan massa. Faucibus. Dolor hendrerit odio nulla.`,
	`Ad ante lectus leo mi tempus nostra sodales venenatis etiam feugiat nullam.`,
	`Vulputate. Lacus at sagittis risus imperdiet parturient auctor interdum. Molestie. Ornare senectus sed.`,
	`Semper blandit porttitor blandit. Id convallis molestie. Litora non posuere dis. Gravida.`,
	`Dolor urna, parturient et nec aliquet suscipit et erat.`,
	`Auctor montes nec ipsum litora felis facilisi blandit convallis quis urna.`,
	`At nascetur ad pharetra potenti vivamus adipiscing erat.`,
	`Porta viverra dictumst taciti accumsan mus penatibus molestie fusce ullamcorper.`,
	`Tellus nostra nascetur cubilia faucibus tempor dignissim. Accumsan ullamcorper, natoque congue.`,
	`Nostra tempus. Elementum dictumst justo, at lorem convallis praesent per.`,
	`Nam. Elit montes. Eu sagittis lobortis conubia integer cubilia montes, nisi erat.`,
	`Porttitor suscipit. Eget, iaculis nullam cras a maecenas ullamcorper.`,
	`Ultricies sit natoque accumsan. Fames primis vitae ante platea.`,
	`Per. Eleifend orci ligula posuere sapien et viverra per morbi aenean sit netus.`,
	`Porta hac. Massa gravida Taciti aptent senectus elementum ullamcorper risus augue.`,
	`Primis vestibulum lacinia vestibulum tellus tortor ligula nascetur.`,
	`Sit justo elit. Hymenaeos sed Vehicula in integer lectus, placerat porttitor, porta lectus.`,
	`Sed inceptos odio penatibus habitant nascetur felis vitae amet scelerisque, suspendisse.`,
	`Consequat pede venenatis accumsan. Ultrices diam fringilla metus hendrerit duis metus.`,
	`Porttitor proin elementum est Aliquet nonummy.`,
	`In. Litora ut. Ultricies mollis curae; potenti velit luctus.`,
	`Placerat Elit curae; hac nullam tempus, feugiat porttitor potenti facilisi quam habitant dolor ornare.`,
	`Nunc per massa eleifend potenti lacus risus odio magna ad, sociis aptent.`,
	`Potenti mi. Ullamcorper. Tincidunt aliquam mollis montes congue. Maecenas magnis aliquet nulla dolor.`,
	`Viverra etiam mattis sociis neque elementum curabitur risus nisl. Fermentum eros mus adipiscing.`,
	`Est, tempus risus aenean platea habitant Curabitur.`,
	`Nulla ridiculus dapibus sagittis ante nunc class arcu etiam, laoreet libero vehicula.`,
	`Curae; duis egestas venenatis phasellus. Nostra vivamus hac.`,
	`Vestibulum amet porttitor ultrices cras curabitur proin. Class lobortis fringilla ac.`,
	`Per rutrum etiam dis. Ultrices amet etiam pretium.`,
	`Lorem mollis dignissim sed maecenas convallis egestas facilisis posuere.`,
	`Laoreet nisi quisque. Sapien nostra, ac dui curabitur pharetra vitae.`,
	`Auctor odio dictum commodo mus est at adipiscing praesent vulputate.`,
	`Augue, turpis aptent egestas. Sollicitudin nullam lectus.`,
	`Potenti egestas at rhoncus viverra neque tristique.`,
	`Euismod nibh quam fringilla volutpat class venenatis eleifend. Hac conubia tempor enim.`,
	`Dolor fusce elementum ridiculus dignissim pretium nulla facilisi.`,
	`Sit vehicula taciti pede facilisi Ad.`,
	`Interdum quisque. Euismod quam urna iaculis. Ultrices Curabitur lacinia donec hymenaeos.`,
	`Et proin fermentum cubilia eros sapien. Habitasse placerat, eu.`,
	`Dui tortor ullamcorper odio vel Taciti pede facilisi nunc.`,
	`Mus leo quisque malesuada sed felis, lectus. Suscipit augue conubia. Suspendisse sagittis sit laoreet.`,
	`Quisque hac sapien purus condimentum urna dictum. Nostra a massa ante non magnis rhoncus.`,
	`Mollis lacinia eget vehicula aliquam massa.`,
	`Etiam venenatis fames. Platea placerat netus consectetuer facilisi suspendisse pede dis.`,
	`Mauris nam, platea non. Scelerisque lacus gravida. Metus bibendum dolor, leo arcu magnis pulvinar.`,
	`Erat cubilia justo commodo consectetuer habitant ut.`,
	`Proin pulvinar blandit eleifend aptent integer libero per non tellus aliquet.`,
	`Eu volutpat. Taciti velit risus nec sociosqu blandit ipsum lectus.`,
	`Duis facilisis metus lorem consectetuer placerat. Cum aptent dapibus.`,
	`In, aliquam vel sapien neque quis. Risus id.`,
	`Elit arcu nunc arcu erat ligula fusce, aliquet tincidunt. Dictum sem sodales aliquet Aliquam.`,
	`Nonummy dignissim cubilia commodo dictum Nisl dapibus orci et quis mollis luctus.`,
	`Enim mauris. Mus velit orci iaculis imperdiet natoque class integer torquent lobortis taciti risus.`,
	`Pharetra. Leo pede natoque adipiscing erat natoque. Placerat nulla ac elementum eu massa.`,
	`Sit. Faucibus leo vulputate torquent pulvinar erat purus fusce varius fames.`,
	`Ad velit imperdiet ridiculus, rhoncus velit. Sociis, sed. Ridiculus commodo.`,
	`Duis hendrerit placerat cubilia interdum. Litora at sodales ultrices elementum rutrum interdum.`,
	`Faucibus consectetuer interdum mi venenatis justo egestas ligula magna.`,
	`Elit vitae. Cras dictum et. Pede. Hendrerit. Imperdiet ad.`,
	`Nisi iaculis. Ut dictum nostra montes orci mollis elit.`,
	`Posuere nascetur. Ridiculus dis quis. Litora torquent lacus mattis quis massa non.`,
	`At nullam quisque sed massa. Lorem at. Adipiscing pede, amet imperdiet fringilla condimentum montes.`,
	`Volutpat nam vulputate convallis class semper.`,
	`Velit morbi, imperdiet, turpis. Eros. Sodales. Mauris. Urna volutpat rhoncus posuere.`,
	`Imperdiet libero facilisis tempus lectus enim molestie. Taciti.`,
	`Lobortis facilisis. Fames mattis phasellus fermentum nostra in donec id lorem porttitor mus sodales.`,
	`Morbi ultrices neque rutrum proin nec.`,
	`Eros, tincidunt orci duis ultricies ad et lacinia hac, dis consectetuer.`,
	`Lacus leo egestas luctus nullam consequat odio amet, eu.`,
}
